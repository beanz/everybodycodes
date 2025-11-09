#!/usr/bin/env perl
use strict;
use warnings;
use v5.30;
use lib "../../lib-perl";
no warnings 'portable';
use Helpers qw/:all/;

sub parse {
  my ($c) = @_;
  chomp($c);
  return split /[:,]/, $c;
}

sub parse2 {
  my ($c) = @_;
  chomp($c);
  return map {[parse($_)]} split /\n/, $c;
}

my @i1 = parse(input(1));
my $a1 = quality(\@i1);
my @i2 = parse2(input(2));
my @q;
for my $l (@i2) {
  push @q, quality($l);
}
@q = sort {$a <=> $b} @q;
my $a2 = $q[$#q] - $q[0];

my @i3 = parse2(input(3));
@q = map {sword($_)} @i3;
@q = sort {compare($b, $a)} @q;
my $a3;
for my $i (0..$#q) {
  $a3 += ($i+1)*$q[$i]->[0];
}

print "Part 1: ", $a1, "\n";
print "Part 2: ", $a2, "\n";
print "Part 3: ", $a3, "\n";

sub compare {
  my ($a, $b) = @_;
  my $f = $a->[1] <=> $b->[1];
  return $f if ($f);
  my $l = min(scalar @{$a->[2]}, scalar @{$b->[2]}) - 1;
  for my $i (0 .. $l) {
    my $f = $a->[2]->[$i] <=> $b->[2]->[$i];
    return $f if ($f);
  }
  return $a->[0] <=> $b->[0]
}

sub sword {
  my ($in) = @_;
  my $id = shift @$in;
  my @fb;
OUTER:
  for my $n (@$in) {
    for my $s (@fb) {
      if ($s->[0] eq "" && $n < $s->[1]) {
        $s->[0] = $n;
        next OUTER;
      }
      if ($s->[2] eq "" && $n > $s->[1]) {
        $s->[2] = $n;
        next OUTER;
      }
    }
    push @fb, ["", $n, ""];
  }
  my $q = join "", map {$_->[1]} @fb;
  my @s = map {join "", @$_} @fb;
  return [$id, $q, \@s];
}

sub quality {
  return sword(@_)->[1];
}
