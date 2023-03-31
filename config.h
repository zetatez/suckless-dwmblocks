
#define UTILS "/home/dionysus/.suckless/suckless-dwmblocks/utils"

//Modify this file to change what commands output to your statusbar, and recompile using the make command.
static const Block blocks[] = {
/*Icon*//*Command*//*Update Interval*//*Update Signal*/
/* {"", UTILS"/status-msg",              3,    0 }, */
{"", UTILS"/status-apps",             1,    0 },
/* {"", UTILS"/status-mail",             300,  0 }, */
{"", UTILS"/status-internet",         1,    0 },
{"", UTILS"/status-battery",          1,    0 },
{"", UTILS"/status-volume",           1,    0 },
{"", UTILS"/status-microphone",       1,    0 },
{"", UTILS"/status-screen-light",     1,    0 },
/* {"", UTILS"/status-cpubar",           3,    0 }, */
{"", UTILS"/status-cpu",              3,    0 },
{"", UTILS"/status-ram",              3,    0 },
{"", UTILS"/status-clock",            1,    0 },
};

//sets delimeter between status commands. NULL character ('\0') means no delimeter.
static char delim[] = " ";
static unsigned int delimLen = 3;
