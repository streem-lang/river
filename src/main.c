#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h> 

#define INSTALL  1
#define REMOVE   2
#define SETUP    3
#define VERSION  4
#define RUN      5
#define BIN      6
const char *RIVER_VERSION = "v0.0+b1";

int type_of_arg (char *arg);

int main(int argc, char *argv[])
{
  if (argc == 1) { argv[1] = "NoArgs"; }
  switch (type_of_arg(argv[1])) {
    case INSTALL:
      if (argc == 3) {
        puts("Downloading project from https://river.heroku.com to temporary folder.....");
        puts("Installing to main loading area....");
        puts("Done!");
        return 0;
      } else if (argc == 2){
        puts("Installing current project....");
        puts("Done!");
        return 0;
      } else {
        puts("Incorrect number of arguments supplied!");
        return 1;
      }
      break;
    case REMOVE:
      if (argc != 3) { puts("Incorrect number of arguments supplied!"); return 1; }
      printf("Removing package %s from main loading area...\n", argv[2]);
      printf("Are you ABSOLUTELY SURE? ");
      char yn = getchar();
      if (yn == 'Y' || yn == 'y') {
        puts("Removing chosen project!");
        puts("Done.");
      } else {
        puts("OK. Exiting now....");
      }
      break;
    case SETUP:
      puts("River doesn't even officially exist yet. This is just a dummy.");
      break;
    case VERSION:
      printf("%s\n", RIVER_VERSION);
      return 0;
      break;
    case RUN:
      puts("There's no Streem interpreter yet!");
      break;
    case BIN:
      puts("There's no Streem interpreter yet!");
      break;
    default:
      puts("River command usage: $ river [command] [arguments ...]");
      puts("COMMANDS: ");
      puts("    install <repo>              Installs package");
      puts("    remove <name>               Uninstalls package from your machene");
      puts("    setup                       Interactive walkthrough setting up your River/Streem project");
      puts("    version                     Shows your your installed River version");
      puts("    run                         Runs your project");
      puts("    bin                         Outputs an executable of your project");
      return 1;
  }
  return 0;
}


int type_of_arg (char *arg)
{
  int type;
  if (strcmp(arg, "install") == 0) {
    type = INSTALL;
  } else if (strcmp(arg, "remove") == 0) {
    type = REMOVE;
  } else if (strcmp(arg, "setup") == 0) {
    type = SETUP;
  } else if (strcmp(arg, "version") == 0) {
    type = VERSION;
  } else if (strcmp(arg, "run") == 0) {
    type = RUN;
  } else if (strcmp(arg, "bin") == 0) {
    type = BIN;
  } else {
    type = 999;
  }

  return type;
}
