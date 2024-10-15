MacOS:
```
brew install gtk+3
brew install pkg-config gtk+3 adwaita-icon-theme
```
Linux:
```
sudo apt update
sudo apt install libgtk-3-dev libcairo2-dev libglib2.0-dev
```
Windows:
```
Install https://chocolatey.org/
PS C:\> choco install golang
PS C:\> choco install git
PS C:\> choco install msys2
PS C:\> mingw64
$ pacman -S mingw-w64-x86_64-gtk3 mingw-w64-x86_64-toolchain base-devel glib2-devel
$ echo 'export PATH=/c/Go/bin:$PATH' >> ~/.bashrc
$ echo 'export PATH=/c/Program\ Files/Git/bin:$PATH' >> ~/.bashrc
$ source ~/.bashrc
$ sed -i -e 's/-Wl,-luuid/-luuid/g' /mingw64/lib/pkgconfig/gdk-3.0.pc # This fixes a bug in pkgconfig
```