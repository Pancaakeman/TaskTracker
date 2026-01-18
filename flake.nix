{
  description = "A simple TaskTracker written in Go (one of the millions on Github)";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = inputs@{ flake-parts, nixpkgs, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin" ];
      perSystem = { pkgs, ... }: {
        packages.default = pkgs.buildGoModule {
          pname = "tasktracker";
          version = "0.1.0";
          src = ./.;
          vendorHash = ""; # Use `nix develop` and run `go mod vendor && nix hash path ./vendor` to fill this in
        };
      };
    };
}
