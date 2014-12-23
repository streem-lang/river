#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#define NORMAL    "\x1B[0m"
#define RED       "\x1B[31m"
#define GREEN     "\x1B[32m"
#define YELLOW    "\x1B[33m"
#define BLUE      "\x1B[34m"
#define MAGNETIA  "\x1B[35m"
#define CYAN      "\x1B[36m"
#define WHITE     "\x1B[37m"

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
        printf("Downloading project from https://river.heroku.com to temporary folder.....\n");
        printf("Installing to main loading area....\n");
        printf("%sDone!%s\n", GREEN, NORMAL);
        return 0;
      } else if (argc == 2){
        printf("Installing current project....\n");
        printf("%sDone!%s\n", GREEN, NORMAL);
        return 0;
      } else {
        printf("Incorrect number of arguments supplied!\n");
        return 1;
      }
      break;
    case REMOVE:
      if (argc != 3) { printf("Incorrect number of arguments supplied!\n"); return 1; }
      printf("Removing package %s from main loading area...\n", argv[2]);
      printf("Are you %sABSOLUTELY SURE%s? ", RED, NORMAL);
      char yn = getchar();
      if (yn == 'Y' || yn == 'y') {
        printf("%sRemoving chosen project.%s\n", RED, NORMAL);
        printf("%sDone!%s\n", GREEN, NORMAL);
      } else {
        printf("%sExiting....%s\n", YELLOW, NORMAL);
      }
      break;
    case SETUP:
      printf("River doesn't even officially exist yet. This is just a dummy.\n");
      printf("%sExiting....%s\n", YELLOW, NORMAL);
      break;
    case VERSION:
      printf("%s\n", RIVER_VERSION);
      return 0;
      break;
    case RUN:
      printf("There's no Streem interpreter yet!\n");
      printf("%sExiting....%s\n", YELLOW, NORMAL);
      break;
    case BIN:
      printf("There's no Streem interpreter yet!\n");
      printf("%sExiting....%s\n", YELLOW, NORMAL);
      break;
    default:
      printf("River command usage: river %s[command]%s [arguments ...]\n", CYAN, NORMAL);
      printf("COMMANDS: \n");
      printf("    %sinstall%s <name>              Installs package\n", CYAN, NORMAL);
      printf("    %sremove%s <name>               Uninstalls package from your machene\n", CYAN, NORMAL);
      printf("    %ssetup%s                       Interactive walkthrough setting up your River/Streem project\n", CYAN, NORMAL);
      printf("    %sversion%s                     Shows your your installed River version\n", CYAN, NORMAL);
      printf("    %srun%s                         Runs your project\n", CYAN, NORMAL);
      printf("    %sbin%s                         Outprintf an executable of your project\n", CYAN, NORMAL);
      printf("%sExiting....%s\n", YELLOW, NORMAL);
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
