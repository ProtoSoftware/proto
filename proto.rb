# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Proto < Formula
  desc "Install and manage Proton versions easily"
  homepage "https://github.com/BitsOfAByte/proto"
  version "0.1.0"
  license "GPL-3.0"
  depends_on :linux

  on_linux do
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.1.0/proto_linux_arm.zip"
      sha256 "b1f95e69d7357dc2d1f6d18613bfafb9ec0b4fe99b3cbc93d10ccb5bb3ecfad9"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.1.0/proto_linux_amd64.zip"
      sha256 "74da2f19eacdf71e900a59e58e19aaf0f98a5c152916657611475ae49e82730d"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.1.0/proto_linux_arm64.zip"
      sha256 "1a2cd229d936a4270d41cc741d965a66c32035634126c399c30c522b9e6e83f8"

      def install
        bin.install "proto"
      end
    end
  end
end
