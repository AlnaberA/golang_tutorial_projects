# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://www.rubydoc.info/github/Homebrew/brew/master/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class GolangTutorialProjects < Formula
  desc "This repository is meant for learning the golang language."
  homepage ""
  url "https://github.com/AlnaberA/golang_tutorial_projects/releases/download/0.0.1/Dance-0.0.1-darwin-amd64.zip"
  sha256 "ec3dac27b566ccbe6216b863451cf6e5ef907856c80edc94896e6f3d6a415547"
  # depends_on "cmake" => :build

  def install
    bin.install "Dance"
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test golang_tutorial_projects`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
