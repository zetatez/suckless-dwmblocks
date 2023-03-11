
#define PWD "/home/dionysus/.suckless/suckless-dwmblocks"

//Modify this file to change what commands output to your statusbar, and recompile using the make command.
static const Block blocks[] = {
/*Icon*//*Command*//*Update Interval*//*Update Signal*/
{"", PWD"/blocks/status-msg",       3,  0 },
{"", PWD"/blocks/status-apps",      1,  0 },
{"", PWD"/blocks/status-mail",      1,  0 },
{"", PWD"/blocks/status-battery",   60, 0 },
{"", PWD"/blocks/status-internet ", 1,  0 },
{"", PWD"/blocks/status-volume",    1,  0 },
/* {"", PWD"/blocks/status-cpubar",    1,  0 }, */
{"", PWD"/blocks/status-cpu",       60, 0 },
{"", PWD"/blocks/status-ram",       15, 0 },
{"", PWD"/blocks/status-clock",     1,  0 },
};

//sets delimeter between status commands. NULL character ('\0') means no delimeter.
static char delim[] = " ";
static unsigned int delimLen = 5;
