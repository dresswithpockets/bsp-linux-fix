# bsp-linux-fix

Patches source engine maps to fix missing models & textures on linux.

Sometimes TF2 maps will present with missing textures when played on linux, but not on windows. This is usually because file system case sensitivity on linux. The mapper packed assets which have non-lowercase characters in their paths, into the map. TF2 doesn't handle these correctly, and won't be able to find the correct model or texture as a result.

## Usage

You must have [`vpkeditcli`](https://github.com/craftablescience/VPKEdit) v4.4 or higher available on path. If you're on linux, your package manager may already have a distribution of it.

Compilation requires Go 1.24 or higher:

```sh
git clone https://github.com/dresswithpockets/bsp-linux-fix.git
cd bsp-linux-fix
go build .
```

Put all the BSPs you wanna patch into a folder, then invoke `bsp-linux-fix` with the input folder path, and any output folder path:

```sh
./bsp-linux-fix inputDir outputDir
```

All patched BSPs will be in the `outputDir` directory. The patched map must exist on any servers you plan on joining. You can play patched maps locally just fine though!

Example output:
```
./bsp-linux-fix tests/in tests/out
Step 1: tests/in/jump_aqua.bsp
Step 1: tests/in/jump_autumn_rc4.bsp
Step 1: tests/in/jump_canyon_b1.bsp
Step 1: tests/in/jump_hexahedron.bsp
Step 1: tests/in/jump_jurf_a2.bsp
Step 1: tests/in/jump_lithium_fix2.bsp
Step 1: tests/in/jump_panama_final.bsp
Step 1: tests/in/jump_quattro_rc1.bsp
Step 1: tests/in/jump_redplanet_v2.bsp
Step 1: tests/in/jump_redplanet_v2_fixed.bsp
Step 1: tests/in/jump_sketchy_final.bsp
Step 1: tests/in/jump_tissue.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_aqua.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_autumn_rc4.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_canyon_b1.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_hexahedron.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_jurf_a2.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_lithium_fix2.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_panama_final.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_quattro_rc1.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_redplanet_v2.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_redplanet_v2_fixed.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_sketchy_final.bsp
Step 2: /tmp/bsplinuxfix_intermediate2616446669/jump_tissue.bsp
```

## Credits

Thanks to https://github.com/spiritov for the idea, and for experimenting on the implementation with me.

Thanks to https://github.com/craftablescience for her tool [VPKEdit](https://github.com/craftablescience/VPKEdit).

Thanks to valve for making a really buggy game.