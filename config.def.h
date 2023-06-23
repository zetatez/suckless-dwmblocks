
#define PREFIX "/home/dionysus/.suckless/suckless-dwmblocks/utils"

//Modify this file to change what commands output to your statusbar, and recompile using the make command.
static const Block blocks[] = {
  /*Icon*//*Command*//*Update Interval*//*Update Signal*/
  // {"", PREFIX"/status-msg"          ,3   ,0 },
  {"", PREFIX"/status-apps"         ,1   ,0 },
  {"", PREFIX"/status-mail"         ,300 ,0 },
  {"", PREFIX"/status-net"          ,1   ,0 },
  {"", PREFIX"/status-battery"      ,1   ,0 },
  {"", PREFIX"/status-volume"       ,1   ,0 },
  {"", PREFIX"/status-micro"        ,1   ,0 },
  {"", PREFIX"/status-screen-light" ,1   ,0 },
  {"", PREFIX"/status-cpu"          ,3   ,0 },
  {"", PREFIX"/status-ram"          ,3   ,0 },
  {"", PREFIX"/status-clock"        ,1   ,0 },
};

//sets delimeter between status commands. NULL character ('\0') means no delimeter.
static char delim[] = "  ";
static unsigned int delimLen = 3;
