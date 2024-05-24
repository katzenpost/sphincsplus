#! /usr/bin/env python3
import itertools

implementations = [
                   ('ref', ['shake', 'sha2', 'haraka']),
                   ('haraka-aesni', ['haraka']),
                   ('shake-avx2', ['shake']),
                   ('shake-a64', ['shake']),
                   ('sha2-avx2', ['sha2']),
                   ]

hashes = ["sha2", "shake", "haraka"]
options = ["f", "s"]
sizes = [128, 192, 256]
thashes = ['robust', 'simple']
platforms = ['linux/arm64', 'linux/amd64', 'darwin/arm64', 'darwin/amd64', 'windows/amd64']

CFLAGS = ['-O3', '-std=c99', '-D CGO=1']
CFLAGS_default = CFLAGS + ['-D HASH=haraka-128',
                           '-D THASH=robust',
                           '-D PARAMS=sphincs-haraka-128f']
CFLAGS_haraka_aesni = CFLAGS + ['-march=native', '-fomit-frame-pointer', '-flto']
CFLAGS_haraka_aesni_default = CFLAGS_haraka_aesni + ['-D HASH=haraka-128',
                                                     '-D THASH=robust',
                                                     '-D PARAMS=sphincs-haraka-128f']
CFLAGS_sha2_avx2 = CFLAGS + ['-march=native', '-fomit-frame-pointer', '-flto']
CFLAGS_sha2_avx2_default = CFLAGS_sha2_avx2 + ['-D HASH=sha2-128',
                                               '-D THASH=robust', 
                                               '-D PARAMS=sphincs-sha2-128f']
CFLAGS_a64 = CFLAGS + ['-march=native', '-fomit-frame-pointer', '-flto']
CFLAGS_a64_default = CFLAGS_a64 + ['-D HASH=shake-128',
                                   '-D THASH=robust',
                                   '-D PARAMS=sphincs-shake-128f'] 
CFLAGS_shake_avx2 = CFLAGS + ['-march=native', '-fomit-frame-pointer', '-flto']
CFLAGS_shake_avx2_default = CFLAGS_shake_avx2 + ['-D HASH=shake-128',
                                                 '-D THASH=robust',
                                                 '-D PARAMS=sphincs-shake-128f']

# Print the default build tags for each platform and implementation specific
# default build tags for each platform
cgo_stanza = set()
for platform in platforms:
    for impl, fns in implementations:
        cflags = None
        x = None
        if impl == 'ref' and (platform.endswith('arm64') or platform.endswith('amd64')):
            x = 'CFLAGS: {}'.format(" ".join(CFLAGS_default))
            cgo_stanza.add("//#cgo {} {}".format(platform, x)) 
            cgo_stanza.add("//#cgo {} test {}".format(platform, x)) 
            cflags = '{} CFLAGS: {}'.format(impl.replace("-", "_"), " ".join(CFLAGS_default))
        elif impl == 'haraka-aesni' and platform.endswith('amd64'):
            cflags = '{} CFLAGS: {}'.format(impl.replace("-", "_"), " ".join(CFLAGS_haraka_aesni_default))
        elif impl == 'shake-avx2' and platform.endswith('amd64'):
            cflags = '{} CFLAGS: {}'.format(impl.replace("-", "_"), " ".join(CFLAGS_shake_avx2_default))
        elif impl == 'sha2-avx2' and platform.endswith('amd64'):
            cflags = '{} CFLAGS: {}'.format(impl.replace("-", "_"), " ".join(CFLAGS_sha2_avx2_default))
        elif impl == 'shake-a64' and platform.endswith('arm64'):
            cflags = '{} CFLAGS: {}'.format(impl.replace("-", "_"), " ".join(CFLAGS_a64_default))
        if cflags is not None:
            cgo_stanza.add("//#cgo {}".format(cflags))

# Print the full permutation of build tags
# default build tags for each platform
full_params = set()
for impl, fns in implementations:
    for fn in fns:
        for opt, size, thash, _hash in itertools.product(options, sizes, thashes, hashes):
            paramset = "sphincs-{}-{}{}".format(fn, size, opt)
            buildtag = "sphincs_{}_{}{} {} {}".format(fn, size, opt, _hash, thash)
            params = "-D PARAMS={}".format(paramset)
            thash = "-D THASH={}".format(thash)
            _hash = "-D HASH={}".format(_hash)
            cflags = 'CFLAGS: {} {} {} {}'.format(" ".join(CFLAGS), _hash, thash, params)
            full_params.add((buildtag,cflags))
for buildtag, flags in full_params:
    cgo_stanza.add("//#cgo {} {}".format(buildtag, flags))

cgo_stanza_sorted = list(cgo_stanza)
cgo_stanza_sorted.sort()
for stanza in cgo_stanza_sorted:
    print(stanza)
