#!/usr/bin/perl
use strict;
use warnings;
use v5.30;
use lib "../../lib-perl";
no warnings 'portable';
use Helpers qw/:all/;

sub parse {
  my ($in) = @_;
  chomp $in;
  $in =~ s/\./0/g;
  my $s = ($in =~ y/#/1/);
  my $m = [map {[split //, $_]} split /\n/, $in];
  my $h = @$m;
  my $w = @{$m->[0]};
  return [$m, $w, $h, $s];
}

my $i1 = parse(input(1));
my $a1 = dig($i1);
my $i2 = parse(input(2));
my $a2 = dig($i2);
my $i3 = parse(input(3));
my $a3 = dig($i3, 1);
print "Part 1: ", $a1, "\n";
print "Part 2: ", $a2, "\n";
print "Part 3: ", $a3, "\n";

sub dig {
  my ($in, $diagonal) = @_;
  my @offsets =
    $diagonal
    ? ([-1, -1], [0, -1], [1, -1], [-1, 0], [1, 0], [-1, 1], [0, 1], [1, 1])
    : ([0, -1], [0, 1], [-1, 0], [1, 0]);
  my $a1 = $in->[3];
  my $done;
  my $n = 2;
  while (!$done) {
    $done = 1;
    for my $y (1 .. $in->[2] - 2) {
    SQ:
      for my $x (1 .. $in->[1] - 2) {
        my $c = $in->[0]->[$y]->[$x];
        if ($c != $n - 1) {
          next;
        }
        for my $o (@offsets) {
          my $c = $in->[0]->[$y + $o->[1]]->[$x + $o->[0]];
          if ($c < $n - 1) {
            next SQ;
          }
        }
        $in->[0]->[$y]->[$x] = $n;
        $a1++;
        $done = 0;
      }
    }
    $n++;
  }
  # my $m = join "", (map {(join "", @$_) . "\n"} @{$in->[0]});
  # print $m;
  return $a1;
}
