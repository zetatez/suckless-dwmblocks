#include<stdlib.h>
#include<stdio.h>
#include<string.h>
#include<unistd.h>
#include<signal.h>

#ifndef NO_X
#include<X11/Xlib.h>
#endif

#ifdef __OpenBSD__
#define SIGPLUS  SIGUSR1+1
#define SIGMINUS SIGUSR1-1
#else
#define SIGPLUS  SIGRTMIN
#define SIGMINUS SIGRTMIN
#endif

#define LENGTH(X)     (sizeof(X) / sizeof (X[0]))
#define CMDLENGTH     50
#define MIN( a, b )   ( ( a < b) ? a : b )
#define STATUSLENGTH  (LENGTH(blocks) * CMDLENGTH + 1)

typedef struct {
  char* icon;
  char* command;
  unsigned int interval;
  unsigned int signal;
} Block;

#ifndef __OpenBSD__
void dummysighandler(int num);
#endif
void sighandler(int num);
void getcmds(int time);
void getsigcmds(unsigned int signal);
void setupsignals();
void sighandler(int signum);
int getstatus(char *str, char *last);
void statusloop();
void termhandler(int signum);
void pstdout();
#ifndef NO_X
void setroot();

static void (*writestatus) () = setroot;
static int setupX();
static Display *dpy;
static int screen;
static Window root;
#else
static void (*writestatus) () = pstdout;
#endif


#include "config.h"

static char statusbar[LENGTH(blocks)][CMDLENGTH] = {0};
static char statusstr[2][STATUSLENGTH];
static int statusContinue = 1;

void getcmd(const Block *block, char *output) {
  size_t icon_len = strlen(block->icon);
  snprintf(output, CMDLENGTH, "%s", block->icon);

  FILE *cmdf = popen(block->command, "r");
  if (!cmdf) return;

  if (fgets(output + icon_len, CMDLENGTH - icon_len - delimLen, cmdf)) {
    size_t len = strlen(output);
    if (len > 0 && output[len - 1] == '\n') {
      output[--len] = '\0';
    }
    if (delim[0] != '\0') {
      strncat(output, delim, CMDLENGTH - len - 1);
    }
  }
  pclose(cmdf);
}

void getcmds(int time)
{
  const Block* current;
  for (unsigned int i = 0; i < LENGTH(blocks); i++) {
    current = blocks + i;
    if ((current->interval != 0 && time % current->interval == 0) || time == -1)
      getcmd(current,statusbar[i]);
  }
}

void getsigcmds(unsigned int signal)
{
  const Block *current;
  for (unsigned int i = 0; i < LENGTH(blocks); i++) {
    current = blocks + i;
    if (current->signal == signal)
      getcmd(current,statusbar[i]);
  }
}

void setupsignals() {
#ifndef __OpenBSD__
  struct sigaction sa = {0};
  sa.sa_handler = dummysighandler;
  for (int i = SIGRTMIN; i <= SIGRTMAX; i++) {
    sigaction(i, &sa, NULL);
  }
#endif

  struct sigaction sa_usr = {0};
  sa_usr.sa_handler = sighandler;
  for (unsigned int i = 0; i < LENGTH(blocks); i++) {
    if (blocks[i].signal > 0) {
      sigaction(SIGMINUS + blocks[i].signal, &sa_usr, NULL);
    }
  }
}

int getstatus(char *str, char *last) {
  str[0] = '\0';
  for (unsigned int i = 0; i < LENGTH(blocks); i++) {
    strncat(str, statusbar[i], STATUSLENGTH - strlen(str) - 1);
  }
  // 删除最后的分隔符
  size_t len = strlen(str);
  if (len >= delimLen) {
    str[len - delimLen] = '\0';
  }
  if (strcmp(str, last) == 0) return 0;
  strcpy(last, str);
  return 1;
}

#ifndef NO_X
void setroot()
{
  if (!getstatus(statusstr[0], statusstr[1])) // Only set root if text has changed.
    return;
  XStoreName(dpy, root, statusstr[0]);
  XFlush(dpy);
}

int setupX()
{
  dpy = XOpenDisplay(NULL);
  if (!dpy) {
    fprintf(stderr, "dwmblocks: Failed to open display\n");
    return 0;
  }
  screen = DefaultScreen(dpy);
  root = RootWindow(dpy, screen);
  return 1;
}
#endif

void pstdout()
{
  if (!getstatus(statusstr[0], statusstr[1])) // Only write out if text has changed.
    return;
  printf("%s\n",statusstr[0]);
  fflush(stdout);
}


void statusloop()
{
  setupsignals();
  int i = 0;
  getcmds(-1);
  while (1) {
    getcmds(i++);
    writestatus();
    if (!statusContinue)
      break;
    sleep(1.0);
  }
}

#ifndef __OpenBSD__
/* this signal handler should do nothing */
void dummysighandler(int signum)
{
  return;
}
#endif

void sighandler(int signum)
{
  getsigcmds(signum-SIGPLUS);
  writestatus();
}

void termhandler(int signum)
{
  statusContinue = 0;
}

int main(int argc, char** argv)
{
  for (int i = 0; i < argc; i++) { // Handle command line arguments
    if (!strcmp("-d",argv[i]))
      strncpy(delim, argv[++i], delimLen);
    else if (!strcmp("-p",argv[i]))
      writestatus = pstdout;
  }
#ifndef NO_X
  if (!setupX())
    return 1;
#endif
  delimLen = MIN(delimLen, strlen(delim));
  delim[delimLen++] = '\0';
  signal(SIGTERM, termhandler);
  signal(SIGINT, termhandler);
  statusloop();
#ifndef NO_X
  XCloseDisplay(dpy);
#endif
  return 0;
}
