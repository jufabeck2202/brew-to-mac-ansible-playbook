package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	output := Parse(fullBrewBundle, emptyAnsibleResult, true)
	if output != fullAnsibleResult {
		t.Errorf("Output was empty")
	}
	fmt.Println(t.TempDir())
}

var fullBrewBundle = `tap "adoptopenjdk/openjdk"
tap "create-go-app/cli"
tap "elastic/tap"
tap "ethereum/ethereum"
tap "hashicorp/tap"
tap "hivemq/mqtt-cli"
tap "homebrew/bundle"
tap "homebrew/cask"
tap "homebrew/cask-drivers"
tap "homebrew/cask-fonts"
tap "homebrew/cask-versions"
tap "homebrew/core"
tap "homebrew/services"
tap "jckuester/tap"
tap "jufabeck2202/formulas", "git@github.com:jufabeck2202/homebrew-formulas.git"
tap "jufabeck2202/tool"
tap "linuxkit/linuxkit"
tap "melonamin/formulae"
tap "nakabonne/ali"
tap "robotsandpencils/made"
tap "tinygo-org/tools"
tap "ubuntu/microk8s"
tap "vapor/tap"
brew "aircrack-ng"
brew "annie"
brew "xz"
brew "ansible"
brew "apache-spark"
brew "node"
brew "apollo-cli"
brew "aria2"
brew "asciinema"
brew "autoconf"
brew "automake"
brew "autossh"
brew "awscli"
brew "azure-cli"
brew "bash-completion"
brew "bat"
brew "bazel", link: false
brew "binutils"
brew "glib"
brew "cairo"
brew "gcc"
brew "gobject-introspection"
brew "harfbuzz"
brew "libraqm"
brew "numpy", link: false
brew "binwalk"
brew "bmon"
brew "boost"
brew "cacli"
brew "carthage"
brew "coreutils"
brew "gnutls"
brew "cask"
brew "suite-sparse"
brew "tbb"
brew "ceres-solver"
brew "cifer"
brew "cloc"
brew "cloud-nuke"
brew "cmake"
brew "crunch"
brew "ctop"
brew "dex2jar"
brew "dns2tcp"
brew "eksctl"
brew "emscripten"
brew "erlang"
brew "esptool", link: false
brew "ethereum"
brew "ruby"
brew "fastlane"
brew "fcrackzip"
brew "libass"
brew "ffmpeg"
brew "fftw"
brew "foremost"
brew "fping"
brew "fswatch"
brew "gcc@8", link: false
brew "gdk-pixbuf"
brew "ghostscript"
brew "git"
brew "glances"
brew "go"
brew "go@1.12"
brew "golang-migrate"
brew "netpbm"
brew "pango"
brew "graphviz"
brew "handbrake"
brew "hashpump"
brew "helm"
brew "htop"
brew "httpie"
brew "hub"
brew "hugo"
brew "libssh"
brew "mysql-client"
brew "hydra"
brew "libheif"
brew "python@3.8"
brew "imagemagick"
brew "istioctl"
brew "john"
brew "jpeg-turbo"
brew "jq"
brew "k3d"
brew "k3sup"
brew "k9s"
brew "kind"
brew "knock"
brew "kompose"
brew "krb5"
brew "kubectx"
brew "latexindent"
brew "libdnet"
brew "libproxy"
brew "libusb"
brew "libwebsockets"
brew "libxml2"
brew "mackup"
brew "mas"
brew "maven"
brew "minikube"
brew "mitmproxy"
brew "qt"
brew "mkvtoolnix"
brew "mosquitto", restart_service: true
brew "mp3wrap"
brew "mp4v2"
brew "vapoursynth"
brew "yt-dlp"
brew "mpv"
brew "mtr"
brew "neovim"
brew "nethogs"
brew "nload"
brew "nmap"
brew "nvm"
brew "pyqt@5"
brew "vtk"
brew "opencv"
brew "opencv@3"
brew "openttd"
brew "operator-sdk"
brew "pdfcpu"
brew "picocom"
brew "pipenv"
brew "pngcheck"
brew "qemu"
brew "podman"
brew "poppler"
brew "postgresql", restart_service: true
brew "pure"
brew "putty"
brew "pygments"
brew "pyqt"
brew "qstat"
brew "rabbitmq", restart_service: true
brew "ranger"
brew "redis", restart_service: true
brew "rename"
brew "ruby@2.5"
brew "rustscan"
brew "rustup-init"
brew "siege"
brew "sip"
brew "skaffold"
brew "socat"
brew "sqlmap"
brew "streamripper"
brew "stress"
brew "subfinder"
brew "swiftgen"
brew "swiftlint"
brew "tcpdump"
brew "tcpflow"
brew "tcpreplay"
brew "tcptrace"
brew "telnet"
brew "the_silver_searcher"
brew "thefuck"
brew "tmux"
brew "tree"
brew "ucspi-tcp"
brew "vips"
brew "wakeonlan"
brew "watch"
brew "wget"
brew "yarn"
brew "youtube-dl"
brew "zsh"
brew "zsh-autosuggestions"
brew "zsh-history-substring-search"
brew "create-go-app/cli/cgapp", link: false
brew "elastic/tap/ecctl"
brew "elastic/tap/metricbeat-full"
brew "hashicorp/tap/terraform"
brew "hivemq/mqtt-cli/mqtt-cli"
brew "jckuester/tap/awsweeper"
brew "linuxkit/linuxkit/linuxkit", args: ["HEAD"]
brew "nakabonne/ali/ali"
brew "robotsandpencils/made/xcodes"
brew "tinygo-org/tools/tinygo"
brew "ubuntu/microk8s/microk8s"
brew "vapor/tap/vapor"
cask "1password"
cask "adobe-acrobat-reader"
cask "adoptopenjdk8"
cask "anaconda", args: { appdir: "/Applications" }
cask "anki", args: { appdir: "/Applications" }
cask "arduino"
cask "arduino-ide-beta"
cask "aria2gui"
cask "balenaetcher"
cask "blender"
cask "burp-suite"
cask "calibre"
cask "cheatsheet"
cask "clipy", args: { appdir: "/Applications" }
cask "dbeaver-community"
cask "discord", args: { appdir: "/Applications" }
cask "docker", args: { appdir: "/Applications" }
cask "dozer", args: { appdir: "/Applications" }
cask "filebot"
cask "fing-cli"
cask "firefox"
cask "flux", args: { appdir: "/Applications" }
cask "folx"
cask "font-source-code-pro"
cask "gimp", args: { appdir: "/Applications" }
cask "google-backup-and-sync", args: { appdir: "/Applications" }
cask "google-chrome", args: { appdir: "/Applications" }
cask "grandperspective"
cask "hyper"
cask "ishowu"
cask "iterm2"
cask "java"
cask "jdownloader", args: { appdir: "/Applications" }
cask "jetbrains-toolbox", args: { appdir: "/Applications" }
cask "macdown"
cask "mactex"
cask "microsoft-auto-update"
cask "microsoft-office"
cask "microsoft-teams"
cask "multipass"
cask "netnewswire"
cask "ngrok", args: { appdir: "/Applications" }
cask "obs"
cask "openscad"
cask "paw"
cask "plex-media-server"
cask "postman"
cask "protege"
cask "pulse"
cask "qlcolorcode"
cask "qlimagesize"
cask "qlmarkdown"
cask "qlstephen"
cask "qlvideo"
cask "quicklook-json"
cask "quicklookase"
cask "redisinsight"
cask "rescuetime"
cask "skype"
cask "snipaste", args: { appdir: "/Applications" }
cask "sourcetree", args: { appdir: "/Applications" }
cask "spotify"
cask "steam"
cask "suspicious-package"
cask "swiftbar"
cask "synology-drive"
cask "tower"
cask "transmission"
cask "ultimaker-cura"
cask "unity-hub"
cask "vlc", args: { appdir: "/Applications" }
cask "vscodium"
cask "webpquicklook"
cask "xtorrent"
cask "zotero"
mas "Actions", id: 1586435171
mas "AdBlock", id: 1402042596
mas "Affinity Designer", id: 824171161
mas "Affinity Photo", id: 824183456
mas "ApolloOne", id: 1044484672
mas "Backtrack", id: 1477089520
mas "Boop", id: 1518425043
mas "Craft", id: 1487937127
mas "DetailsPro", id: 1524366536
mas "DevCleaner", id: 1388020431
mas "Developer", id: 640199958
mas "Disk Speed Test", id: 425264550
mas "Elmedia Video Player", id: 1044549675
mas "Emporter", id: 1406832001
mas "Gifski", id: 1351639930
mas "GoodNotes", id: 1444383602
mas "HEIC Converter", id: 1294126402
mas "HLTH", id: 1503879351
mas "Icon Set Creator", id: 939343785
mas "Judo", id: 1564578427
mas "Keynote", id: 409183694
mas "Magnet", id: 441258766
mas "Markdown Editor", id: 1458220908
mas "Microsoft Remote Desktop", id: 1295203466
mas "MindNode", id: 1289197285
mas "Network Speed Tester", id: 1217419133
mas "NordVPN", id: 905953485
mas "NordVPN IKE", id: 1116599239
mas "Notability", id: 360593530
mas "Numbers", id: 409203825
mas "Pages", id: 409201541
mas "PDF Merge & PDF Splitter +", id: 651952889
mas "Ping Info", id: 1505603580
mas "Plash", id: 1494023538
mas "Playgrounds", id: 1496833156
mas "RocketSim", id: 1504940162
mas "SpoticaMenu", id: 570549457
mas "The Unarchiver", id: 425424353
mas "ToothFairy", id: 1191449274
mas "Tot", id: 1491071483
mas "Video Joiner & Merger", id: 1349784180
mas "XCOrganizer", id: 1507556912
mas "Yoink", id: 457622435
`

var emptyAnsibleResult = `---
downloads: ~/.ansible-downloads/

configure_dotfiles: true
configure_terminal: true
configure_osx: true

# Set to 'true' to configure the Dock via dockutil.
configure_dock: false
dockitems_remove: []
# - Launchpad
# - TV
# - Podcasts
# - 'App Store'
dockitems_persist: []
# - name: "Sublime Text"
#   path: "/Applications/Sublime Text.app/"
#   pos: 5

configure_sudoers: false
sudoers_custom_config: ''
# Example:
# sudoers_custom_config: |
#   # Allow users in admin group to use sudo with no password.
#   %admin ALL=(ALL) NOPASSWD: ALL

dotfiles_repo: https://github.com/geerlingguy/dotfiles.git
dotfiles_repo_accept_hostkey: true
dotfiles_repo_local_destination: ~/Development/GitHub/dotfiles
dotfiles_files:
  - .zshrc
  - .gitignore
  - .inputrc
  - .osx
  - .vimrc

homebrew_installed_packages:
homebrew_taps:
homebrew_cask_appdir: /Applications
homebrew_cask_apps:
mas_installed_apps: []
mas_email: ""
mas_password: ""

osx_script: "~/.osx --no-restart"

# Install packages from other package managers.
# Note: You are responsible for making sure the required package managers are
# installed, eg. through homebrew.
composer_packages: []
# - name: drush
#   state: present # present/absent, default: present
#   version: "^8.1" # default: N/A
gem_packages: []
# - name: bundler
#   state: present # present/absent/latest, default: present
#   version: "~> 1.15.1" # default: N/A
npm_packages: []
# - name: webpack
#   state: present # present/absent/latest, default: present
#   version: "^2.6" # default: N/A
pip_packages: []
# - name: mkdocs
#   state: present # present/absent/latest, default: present
#   version: "0.16.3" # default: N/A

# Set to 'true' to configure Sublime Text.
configure_sublime: false
sublime_base_path: '~/Library/Application Support/Sublime Text'
sublime_config_path: "Packages/User"
sublime_package_control:
  - "DocBlockr"
  - "Dockerfile Syntax Highlighting"
  - "FileDiffs"
  - "GitHub Flavored Markdown Preview"
  - "Jinja2"
  - "Package Control"
  - "PHP-Twig"
  - "Pretty JSON"
  - "SublimeLinter"
  - "SublimeLinter-contrib-yamllint"
  - "TrailingSpaces"
  - "WordCount"

# Glob pattern to ansible task files to run after all other tasks are finished.
post_provision_tasks: []
`
var fullAnsibleResult = `---
downloads: ~/.ansible-downloads/

configure_dotfiles: true
configure_terminal: true
configure_osx: true

# Set to 'true' to configure the Dock via dockutil.
configure_dock: false
dockitems_remove: []
# - Launchpad
# - TV
# - Podcasts
# - 'App Store'
dockitems_persist: []
# - name: "Sublime Text"
#   path: "/Applications/Sublime Text.app/"
#   pos: 5

configure_sudoers: false
sudoers_custom_config: ''
# Example:
# sudoers_custom_config: |
#   # Allow users in admin group to use sudo with no password.
#   %admin ALL=(ALL) NOPASSWD: ALL

dotfiles_repo: https://github.com/geerlingguy/dotfiles.git
dotfiles_repo_accept_hostkey: true
dotfiles_repo_local_destination: ~/Development/GitHub/dotfiles
dotfiles_files:
  - .zshrc
  - .gitignore
  - .inputrc
  - .osx
  - .vimrc

homebrew_installed_packages:
  - aircrack-ng    # added 2021-12-17
  - annie    # added 2021-12-17
  - xz    # added 2021-12-17
  - ansible    # added 2021-12-17
  - apache-spark    # added 2021-12-17
  - node    # added 2021-12-17
  - apollo-cli    # added 2021-12-17
  - aria2    # added 2021-12-17
  - asciinema    # added 2021-12-17
  - autoconf    # added 2021-12-17
  - automake    # added 2021-12-17
  - autossh    # added 2021-12-17
  - awscli    # added 2021-12-17
  - azure-cli    # added 2021-12-17
  - bash-completion    # added 2021-12-17
  - bat    # added 2021-12-17
  - bazel    # added 2021-12-17
  - binutils    # added 2021-12-17
  - glib    # added 2021-12-17
  - cairo    # added 2021-12-17
  - gcc    # added 2021-12-17
  - gobject-introspection    # added 2021-12-17
  - harfbuzz    # added 2021-12-17
  - libraqm    # added 2021-12-17
  - numpy    # added 2021-12-17
  - binwalk    # added 2021-12-17
  - bmon    # added 2021-12-17
  - boost    # added 2021-12-17
  - cacli    # added 2021-12-17
  - carthage    # added 2021-12-17
  - coreutils    # added 2021-12-17
  - gnutls    # added 2021-12-17
  - cask    # added 2021-12-17
  - suite-sparse    # added 2021-12-17
  - tbb    # added 2021-12-17
  - ceres-solver    # added 2021-12-17
  - cifer    # added 2021-12-17
  - cloc    # added 2021-12-17
  - cloud-nuke    # added 2021-12-17
  - cmake    # added 2021-12-17
  - crunch    # added 2021-12-17
  - ctop    # added 2021-12-17
  - dex2jar    # added 2021-12-17
  - dns2tcp    # added 2021-12-17
  - eksctl    # added 2021-12-17
  - emscripten    # added 2021-12-17
  - erlang    # added 2021-12-17
  - esptool    # added 2021-12-17
  - ethereum    # added 2021-12-17
  - ruby    # added 2021-12-17
  - fastlane    # added 2021-12-17
  - fcrackzip    # added 2021-12-17
  - libass    # added 2021-12-17
  - ffmpeg    # added 2021-12-17
  - fftw    # added 2021-12-17
  - foremost    # added 2021-12-17
  - fping    # added 2021-12-17
  - fswatch    # added 2021-12-17
  - gcc@8    # added 2021-12-17
  - gdk-pixbuf    # added 2021-12-17
  - ghostscript    # added 2021-12-17
  - git    # added 2021-12-17
  - glances    # added 2021-12-17
  - go    # added 2021-12-17
  - go@1.12    # added 2021-12-17
  - golang-migrate    # added 2021-12-17
  - netpbm    # added 2021-12-17
  - pango    # added 2021-12-17
  - graphviz    # added 2021-12-17
  - handbrake    # added 2021-12-17
  - hashpump    # added 2021-12-17
  - helm    # added 2021-12-17
  - htop    # added 2021-12-17
  - httpie    # added 2021-12-17
  - hub    # added 2021-12-17
  - hugo    # added 2021-12-17
  - libssh    # added 2021-12-17
  - mysql-client    # added 2021-12-17
  - hydra    # added 2021-12-17
  - libheif    # added 2021-12-17
  - python@3.8    # added 2021-12-17
  - imagemagick    # added 2021-12-17
  - istioctl    # added 2021-12-17
  - john    # added 2021-12-17
  - jpeg-turbo    # added 2021-12-17
  - jq    # added 2021-12-17
  - k3d    # added 2021-12-17
  - k3sup    # added 2021-12-17
  - k9s    # added 2021-12-17
  - kind    # added 2021-12-17
  - knock    # added 2021-12-17
  - kompose    # added 2021-12-17
  - krb5    # added 2021-12-17
  - kubectx    # added 2021-12-17
  - latexindent    # added 2021-12-17
  - libdnet    # added 2021-12-17
  - libproxy    # added 2021-12-17
  - libusb    # added 2021-12-17
  - libwebsockets    # added 2021-12-17
  - libxml2    # added 2021-12-17
  - mackup    # added 2021-12-17
  - mas    # added 2021-12-17
  - maven    # added 2021-12-17
  - minikube    # added 2021-12-17
  - mitmproxy    # added 2021-12-17
  - qt    # added 2021-12-17
  - mkvtoolnix    # added 2021-12-17
  - mosquitto    # added 2021-12-17
  - mp3wrap    # added 2021-12-17
  - mp4v2    # added 2021-12-17
  - vapoursynth    # added 2021-12-17
  - yt-dlp    # added 2021-12-17
  - mpv    # added 2021-12-17
  - mtr    # added 2021-12-17
  - neovim    # added 2021-12-17
  - nethogs    # added 2021-12-17
  - nload    # added 2021-12-17
  - nmap    # added 2021-12-17
  - nvm    # added 2021-12-17
  - pyqt@5    # added 2021-12-17
  - vtk    # added 2021-12-17
  - opencv    # added 2021-12-17
  - opencv@3    # added 2021-12-17
  - openttd    # added 2021-12-17
  - operator-sdk    # added 2021-12-17
  - pdfcpu    # added 2021-12-17
  - picocom    # added 2021-12-17
  - pipenv    # added 2021-12-17
  - pngcheck    # added 2021-12-17
  - qemu    # added 2021-12-17
  - podman    # added 2021-12-17
  - poppler    # added 2021-12-17
  - postgresql    # added 2021-12-17
  - pure    # added 2021-12-17
  - putty    # added 2021-12-17
  - pygments    # added 2021-12-17
  - pyqt    # added 2021-12-17
  - qstat    # added 2021-12-17
  - rabbitmq    # added 2021-12-17
  - ranger    # added 2021-12-17
  - redis    # added 2021-12-17
  - rename    # added 2021-12-17
  - ruby@2.5    # added 2021-12-17
  - rustscan    # added 2021-12-17
  - rustup-init    # added 2021-12-17
  - siege    # added 2021-12-17
  - sip    # added 2021-12-17
  - skaffold    # added 2021-12-17
  - socat    # added 2021-12-17
  - sqlmap    # added 2021-12-17
  - streamripper    # added 2021-12-17
  - stress    # added 2021-12-17
  - subfinder    # added 2021-12-17
  - swiftgen    # added 2021-12-17
  - swiftlint    # added 2021-12-17
  - tcpdump    # added 2021-12-17
  - tcpflow    # added 2021-12-17
  - tcpreplay    # added 2021-12-17
  - tcptrace    # added 2021-12-17
  - telnet    # added 2021-12-17
  - the_silver_searcher    # added 2021-12-17
  - thefuck    # added 2021-12-17
  - tmux    # added 2021-12-17
  - tree    # added 2021-12-17
  - ucspi-tcp    # added 2021-12-17
  - vips    # added 2021-12-17
  - wakeonlan    # added 2021-12-17
  - watch    # added 2021-12-17
  - wget    # added 2021-12-17
  - yarn    # added 2021-12-17
  - youtube-dl    # added 2021-12-17
  - zsh    # added 2021-12-17
  - zsh-autosuggestions    # added 2021-12-17
  - zsh-history-substring-search    # added 2021-12-17
  - create-go-app/cli/cgapp    # added 2021-12-17
  - elastic/tap/ecctl    # added 2021-12-17
  - elastic/tap/metricbeat-full    # added 2021-12-17
  - hashicorp/tap/terraform    # added 2021-12-17
  - hivemq/mqtt-cli/mqtt-cli    # added 2021-12-17
  - jckuester/tap/awsweeper    # added 2021-12-17
  - linuxkit/linuxkit/linuxkit    # added 2021-12-17
  - nakabonne/ali/ali    # added 2021-12-17
  - robotsandpencils/made/xcodes    # added 2021-12-17
  - tinygo-org/tools/tinygo    # added 2021-12-17
  - ubuntu/microk8s/microk8s    # added 2021-12-17
  - vapor/tap/vapor    # added 2021-12-17
homebrew_taps:
  - adoptopenjdk/openjdk    # added 2021-12-17
  - create-go-app/cli    # added 2021-12-17
  - elastic/tap    # added 2021-12-17
  - ethereum/ethereum    # added 2021-12-17
  - hashicorp/tap    # added 2021-12-17
  - hivemq/mqtt-cli    # added 2021-12-17
  - homebrew/bundle    # added 2021-12-17
  - homebrew/cask    # added 2021-12-17
  - homebrew/cask-drivers    # added 2021-12-17
  - homebrew/cask-fonts    # added 2021-12-17
  - homebrew/cask-versions    # added 2021-12-17
  - homebrew/core    # added 2021-12-17
  - homebrew/services    # added 2021-12-17
  - jckuester/tap    # added 2021-12-17
  - jufabeck2202/formulas    # added 2021-12-17
  - jufabeck2202/tool    # added 2021-12-17
  - linuxkit/linuxkit    # added 2021-12-17
  - melonamin/formulae    # added 2021-12-17
  - nakabonne/ali    # added 2021-12-17
  - robotsandpencils/made    # added 2021-12-17
  - tinygo-org/tools    # added 2021-12-17
  - ubuntu/microk8s    # added 2021-12-17
  - vapor/tap    # added 2021-12-17
homebrew_cask_appdir: /Applications
homebrew_cask_apps:
  - 1password    # added 2021-12-17
  - adobe-acrobat-reader    # added 2021-12-17
  - adoptopenjdk8    # added 2021-12-17
  - anaconda    # added 2021-12-17
  - anki    # added 2021-12-17
  - arduino    # added 2021-12-17
  - arduino-ide-beta    # added 2021-12-17
  - aria2gui    # added 2021-12-17
  - balenaetcher    # added 2021-12-17
  - blender    # added 2021-12-17
  - burp-suite    # added 2021-12-17
  - calibre    # added 2021-12-17
  - cheatsheet    # added 2021-12-17
  - clipy    # added 2021-12-17
  - dbeaver-community    # added 2021-12-17
  - discord    # added 2021-12-17
  - docker    # added 2021-12-17
  - dozer    # added 2021-12-17
  - filebot    # added 2021-12-17
  - fing-cli    # added 2021-12-17
  - firefox    # added 2021-12-17
  - flux    # added 2021-12-17
  - folx    # added 2021-12-17
  - font-source-code-pro    # added 2021-12-17
  - gimp    # added 2021-12-17
  - google-backup-and-sync    # added 2021-12-17
  - google-chrome    # added 2021-12-17
  - grandperspective    # added 2021-12-17
  - hyper    # added 2021-12-17
  - ishowu    # added 2021-12-17
  - iterm2    # added 2021-12-17
  - java    # added 2021-12-17
  - jdownloader    # added 2021-12-17
  - jetbrains-toolbox    # added 2021-12-17
  - macdown    # added 2021-12-17
  - mactex    # added 2021-12-17
  - microsoft-auto-update    # added 2021-12-17
  - microsoft-office    # added 2021-12-17
  - microsoft-teams    # added 2021-12-17
  - multipass    # added 2021-12-17
  - netnewswire    # added 2021-12-17
  - ngrok    # added 2021-12-17
  - obs    # added 2021-12-17
  - openscad    # added 2021-12-17
  - paw    # added 2021-12-17
  - plex-media-server    # added 2021-12-17
  - postman    # added 2021-12-17
  - protege    # added 2021-12-17
  - pulse    # added 2021-12-17
  - qlcolorcode    # added 2021-12-17
  - qlimagesize    # added 2021-12-17
  - qlmarkdown    # added 2021-12-17
  - qlstephen    # added 2021-12-17
  - qlvideo    # added 2021-12-17
  - quicklook-json    # added 2021-12-17
  - quicklookase    # added 2021-12-17
  - redisinsight    # added 2021-12-17
  - rescuetime    # added 2021-12-17
  - skype    # added 2021-12-17
  - snipaste    # added 2021-12-17
  - sourcetree    # added 2021-12-17
  - spotify    # added 2021-12-17
  - steam    # added 2021-12-17
  - suspicious-package    # added 2021-12-17
  - swiftbar    # added 2021-12-17
  - synology-drive    # added 2021-12-17
  - tower    # added 2021-12-17
  - transmission    # added 2021-12-17
  - ultimaker-cura    # added 2021-12-17
  - unity-hub    # added 2021-12-17
  - vlc    # added 2021-12-17
  - vscodium    # added 2021-12-17
  - webpquicklook    # added 2021-12-17
  - xtorrent    # added 2021-12-17
  - zotero    # added 2021-12-17
mas_installed_apps: []
  - 1586435171    # added (Actions) 2021-12-17
  - 1402042596    # added (AdBlock) 2021-12-17
  - 824171161    # added (Affinity Designer) 2021-12-17
  - 824183456    # added (Affinity Photo) 2021-12-17
  - 1044484672    # added (ApolloOne) 2021-12-17
  - 1477089520    # added (Backtrack) 2021-12-17
  - 1518425043    # added (Boop) 2021-12-17
  - 1487937127    # added (Craft) 2021-12-17
  - 1524366536    # added (DetailsPro) 2021-12-17
  - 1388020431    # added (DevCleaner) 2021-12-17
  - 640199958    # added (Developer) 2021-12-17
  - 425264550    # added (Disk Speed Test) 2021-12-17
  - 1044549675    # added (Elmedia Video Player) 2021-12-17
  - 1406832001    # added (Emporter) 2021-12-17
  - 1351639930    # added (Gifski) 2021-12-17
  - 1444383602    # added (GoodNotes) 2021-12-17
  - 1294126402    # added (HEIC Converter) 2021-12-17
  - 1503879351    # added (HLTH) 2021-12-17
  - 939343785    # added (Icon Set Creator) 2021-12-17
  - 1564578427    # added (Judo) 2021-12-17
  - 409183694    # added (Keynote) 2021-12-17
  - 441258766    # added (Magnet) 2021-12-17
  - 1458220908    # added (Markdown Editor) 2021-12-17
  - 1295203466    # added (Microsoft Remote Desktop) 2021-12-17
  - 1289197285    # added (MindNode) 2021-12-17
  - 1217419133    # added (Network Speed Tester) 2021-12-17
  - 905953485    # added (NordVPN) 2021-12-17
  - 1116599239    # added (NordVPN IKE) 2021-12-17
  - 360593530    # added (Notability) 2021-12-17
  - 409203825    # added (Numbers) 2021-12-17
  - 409201541    # added (Pages) 2021-12-17
  - 651952889    # added (PDF Merge & PDF Splitter +) 2021-12-17
  - 1505603580    # added (Ping Info) 2021-12-17
  - 1494023538    # added (Plash) 2021-12-17
  - 1496833156    # added (Playgrounds) 2021-12-17
  - 1504940162    # added (RocketSim) 2021-12-17
  - 570549457    # added (SpoticaMenu) 2021-12-17
  - 425424353    # added (The Unarchiver) 2021-12-17
  - 1191449274    # added (ToothFairy) 2021-12-17
  - 1491071483    # added (Tot) 2021-12-17
  - 1349784180    # added (Video Joiner & Merger) 2021-12-17
  - 1507556912    # added (XCOrganizer) 2021-12-17
  - 457622435    # added (Yoink) 2021-12-17
mas_email: ""
mas_password: ""

osx_script: "~/.osx --no-restart"

# Install packages from other package managers.
# Note: You are responsible for making sure the required package managers are
# installed, eg. through homebrew.
composer_packages: []
# - name: drush
#   state: present # present/absent, default: present
#   version: "^8.1" # default: N/A
gem_packages: []
# - name: bundler
#   state: present # present/absent/latest, default: present
#   version: "~> 1.15.1" # default: N/A
npm_packages: []
# - name: webpack
#   state: present # present/absent/latest, default: present
#   version: "^2.6" # default: N/A
pip_packages: []
# - name: mkdocs
#   state: present # present/absent/latest, default: present
#   version: "0.16.3" # default: N/A

# Set to 'true' to configure Sublime Text.
configure_sublime: false
sublime_base_path: '~/Library/Application Support/Sublime Text'
sublime_config_path: "Packages/User"
sublime_package_control:
  - "DocBlockr"
  - "Dockerfile Syntax Highlighting"
  - "FileDiffs"
  - "GitHub Flavored Markdown Preview"
  - "Jinja2"
  - "Package Control"
  - "PHP-Twig"
  - "Pretty JSON"
  - "SublimeLinter"
  - "SublimeLinter-contrib-yamllint"
  - "TrailingSpaces"
  - "WordCount"

# Glob pattern to ansible task files to run after all other tasks are finished.
post_provision_tasks: []
`
