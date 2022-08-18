# stats

Shortlink stats-server

### Build

```bash
$> mkdir build && cd build
$> conan install ..
$> cmake .. -DCMAKE_EXPORT_COMPILE_COMMANDS=1 # generates compile_commands.json
$> ln -s compile_commands.json ../compile_commands.json # link compile_commands.json to home dir
```