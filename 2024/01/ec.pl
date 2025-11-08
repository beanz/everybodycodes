#!/usr/bin/perl
use strict;
use warnings;
use v5.30;
use lib "../../lib-perl";
no warnings 'portable';
use Helpers qw/:all/;

my $i1 = input(1);
my $a1 = 3 * ($i1 =~ y/C//) + ($i1 =~ y/B//);
my $i2 = input(2);
my @i2 = $i2 =~ m!(..)!g;
my $a2 = 0;
for my $pair (@i2) {
  if ($pair eq "xx") {
    next;
  }
  if ($pair =~ s/x//g) {
    $a2 += {"A" => 0, "B" => 1, "C" => 3, "D" => 5}->{$pair};
    next;
  }
  $a2 +=
    6 * ($pair =~ y/D//) +
    4 * ($pair =~ y/C//) +
    2 * ($pair =~ y/B//) +
    ($pair =~ y/A//);
}

my $i3 = input(3);
my @i3 = ($i3 =~ m!(...)!g);
my $a3 = potion3(\@i3);
print "Part 1: ", $a1, "\n";
print "Part 2: ", $a2, "\n";
print "Part 3: ", $a3, "\n";

sub potion3 {
  my ($in) = @_;
  my $a = 0;
  for my $p (@$in) {
    my $xs = ($p =~ y/x/X/);
    my $nx = (length $p) - $xs;
    if ($nx == 0) {
      next;
    }
    my $extra = 2 - $xs;
    my $n =
      (5 + $extra) * ($p =~ y/D//) +
      (3 + $extra) * ($p =~ y/C//) +
      (1 + $extra) * ($p =~ y/B//) +
      (0 + $extra) * ($p =~ y/A//);
    $a += $n;
  }
  return $a;
}
