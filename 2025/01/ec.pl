#!/usr/bin/perl
use strict;
use warnings;
use v5.30;
use lib "../../lib-perl";
no warnings 'portable';
use Helpers qw/:all/;

sub parse {
  my ($c) = @_;
  chomp($c);
  my ($n, $m) = split /\n\n/, $c;
  my @names = split /,/, $n;
  $m =~ y/LR/-+/;
  my @moves = split /,/, $m;
  return \@names, \@moves;
}

my ($names1, $moves1) = parse(input(1));
my $p1 = 0;
for my $m (@$moves1) {
  $p1 += $m;
  if ($p1 >= @$names1) {
    $p1 = @$names1-1;
  } elsif ($p1 < 0) {
    $p1 = 0;
  }
}
my $a1 = $names1->[$p1];
my ($names2, $moves2) = parse(input(2));
my $p2 = 0;
for my $m (@$moves2) {
  $p2 += $m;
  $p2 %= @$names2;
}
my $a2 = $names2->[$p2];
my ($names3, $moves3) = parse(input(3));
for my $m (@$moves3) {
  my $p3 = $m%@$names3;
  ($names3->[0], $names3->[$p3]) = ($names3->[$p3], $names3->[0]);
}
my $a3 = $names3->[0];

print "Part 1: ", $a1, "\n";
print "Part 2: ", $a2, "\n";
print "Part 3: ", $a3, "\n";
