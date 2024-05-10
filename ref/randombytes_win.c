//go:build cgo && (windows && amd64)

/*
This code was derived from highctidh and is public domain.
*/
#include "randombytes_win.h"

ssize_t getrandom(char *buf, size_t buflen);
#include <windows.h>
#include <ntsecapi.h>
#define getrandom(x, y) RtlGenRandom(x, y)

void randombytes(unsigned char * x, unsigned long long xlen)
{
  ssize_t n;
  n = getrandom((char *) x, xlen);
  if (n != 1) {
      exit(2);
  }
}
