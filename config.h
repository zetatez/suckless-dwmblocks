
#define PREFIX "/home/dionysus/.suckless/suckless-dwmblocks/cmds/bins"

//Modify this file to change what commands output to your statusbar, and recompile using the make command.
static const Block blocks[] = {
  /*
   * Icon, Command, Update Interval, Update Signal
  */
  {"", PREFIX"/msg"          ,5     ,0 },
  {"", PREFIX"/apps"         ,8     ,0 },
  {"", PREFIX"/mail"         ,900   ,0 },
  {"", PREFIX"/net"          ,1     ,0 },
  {"", PREFIX"/battery"      ,180   ,0 },
  {"", PREFIX"/volume"       ,1     ,0 },
  {"", PREFIX"/micro"        ,1     ,0 },
  {"", PREFIX"/screen-light" ,1     ,0 },
  {"", PREFIX"/cpu"          ,3     ,0 },
  {"", PREFIX"/ram"          ,3     ,0 },
  {"", PREFIX"/weather"      ,3600  ,0 },
  {"", PREFIX"/clock"        ,1     ,0 },
};

//sets delimeter between status commands. NULL character ('\0') means no delimeter.
static char delim[] = " ";
static unsigned int delimLen = 3;
