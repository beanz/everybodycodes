#!/usr/bin/perl
use strict;
use warnings;
use v5.30;
use lib "../../lib-perl";
no warnings 'portable';
use Helpers qw/:all/;

my $a1= potions(input(1), 1);
my $a2 = potions(input(2), 2);
my $a3 = potions(input(3), 3);
print "Part 1: ", $a1, "\n";
print "Part 2: ", $a2, "\n";
print "Part 3: ", $a3, "\n";

sub potions {
  my ($in, $n) = @_;
  my @in = unpack "(A$n)*", $in;
  my $a = 0;
  for my $p (@in) {
    my $xs = ($p =~ y/x/X/);
    my $nx = (length $p) - $xs;
    if ($nx == 0) {
      next;
    }
    my $extra = $nx - 1;
    my $n =
      (5 + $extra) * ($p =~ y/D//) +
      (3 + $extra) * ($p =~ y/C//) +
      (1 + $extra) * ($p =~ y/B//) +
      (0 + $extra) * ($p =~ y/A//);
    $a += $n;
  }
  return $a;
}
