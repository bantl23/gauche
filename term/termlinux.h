#include <unistd.h>
#include <termios.h>

struct termios raw;
struct termios saved;

int restore()
{
    return tcsetattr(STDIN_FILENO, TCSANOW, &saved);
}

int init()
{
    int rv = tcgetattr(STDIN_FILENO, &saved);
    if (rv != 0)
    {
        return rv;
    }

    raw = saved;
    raw.c_iflag &= ~(BRKINT | ICRNL | INPCK | ISTRIP | IXON);
    raw.c_oflag &= ~(OPOST);
    raw.c_cflag |= (CS8);
    raw.c_lflag &= ~(ECHO | ICANON | IEXTEN | ISIG);
    raw.c_cc[VMIN] = 0;
    raw.c_cc[VTIME] = 1;

    return tcsetattr(STDIN_FILENO, TCSAFLUSH, &raw);
}
