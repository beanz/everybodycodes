package Helpers;
use strict;
use warnings;
use v5.30;
use Exporter qw(import);

use constant {
  DEBUG => $ENV{EC_DEBUG} // 0,
  TEST => $ENV{EC_TEST} // 0,
};

use Storable qw/dclone/;
use List::Util qw/min max minstr maxstr sum pairs product uniq reduce/;
use List::MoreUtils qw/zip pairwise minmax/;
use POSIX qw/ceil floor round/;

our %EXPORT_TAGS = (
  'all' => [
    qw(
      DEBUG
      TEST
      dclone
      min max minstr maxstr sum pairs product uniq reduce
      zip pairwise minmax
      ceil floor round

      slurp
      input_file
      input
    )
  ],
);
our @EXPORT_OK = (@{$EXPORT_TAGS{'all'}});
our $VERSION = '0.01';

use FindBin;

sub slurp {
  my ($f) = @_;
  open my $fh, '<', $f or die "reading $f: $!";
  local $/;
  my $c = <$fh>;
  close $fh;
  return $c
}

sub input_file {
  my ($part, $ex) = @_;
  my ($y, $d) = ($FindBin::RealBin =~ m!(\d{4})/(\d{2})!);
  my $path = sprintf "$FindBin::RealBin/../../input/$y/$d";
  if ($ex) {
    return sprintf "%s/%s-p%d.txt", $path, $ex, $part;
  }
  return sprintf "%s/everybody_codes_e%d_q%02d_p%d.txt", $path, $y, $d, $part;
}

sub input {
  my $f = input_file(@_);
  return slurp($f);
}

1;
