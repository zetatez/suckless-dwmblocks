
#define PREFIX "/home/dionysus/git/suckless-dwmblocks/cmds/bins"

//Modify this file to change what commands output to your statusbar, and recompile using the make command.
static const Block blocks[] = {
  /*
   * Icon, Command, Update Interval, Update Signal
  */
  {"", PREFIX"/clean-msg"    ,30    ,0 },
  {"", PREFIX"/msg"          ,1     ,0 },
  {"", PREFIX"/procs"        ,3     ,0 },
  {"", PREFIX"/email"        ,900   ,0 },
  {"", PREFIX"/net"          ,3     ,0 },
  {"", PREFIX"/battery"      ,180   ,0 },
  {"", PREFIX"/volume"       ,1     ,0 },
  {"", PREFIX"/micro"        ,1     ,0 },
  {"", PREFIX"/screen-light" ,1     ,0 },
  {"", PREFIX"/cpu"          ,3     ,0 },
  {"", PREFIX"/ram"          ,3     ,0 },
  // {"", PREFIX"/disk"         ,60    ,0 },
  // {"", PREFIX"/weather"      ,3600  ,0 },
  {"", PREFIX"/clock"        ,1     ,0 },
};

//sets delimeter between status commands. NULL character ('\0') means no delimeter.
static char delim[] = "  ";
static unsigned int delimLen = 3;
