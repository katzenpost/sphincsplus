#ifndef SPX_RANDOMBYTES_WIN_H
#define SPX_RANDOMBYTES_WIN_H

extern void randombytes(unsigned char * x,unsigned long long xlen);
#include <basetsd.h>
#define ssize_t SSIZE_T
#include <windows.h>
#define SystemFunction036 NTAPI SystemFunction036
#include <ntsecapi.h>
#undef SystemFunction036
#pragma comment(lib, "advapi32.lib")

#endif
