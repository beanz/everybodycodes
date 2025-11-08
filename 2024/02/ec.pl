#!/usr/bin/perl
use strict;
use warnings;
use v5.30;
use lib "../../lib-perl";
no warnings 'portable';
use Helpers qw/:all/;

my $i1 = input(1);
my $a1 = 0;
{
  my ($w, $s) = split /\n\n/, $i1;
  $w =~ s/WORDS://;
  my @w = split /,/, $w;
  for my $w (@w) {
    $a1 += ($s =~ s/$w/$w/g);
  }
}

my $i2 = input(2);
my $a2 = 0;
{
  chomp $i2;
  my ($w2, $ss) = split /\n\n/, $i2;
  $w2 =~ s/WORDS://;
  my @w = split /,/, $w2;
  for my $s (split /\n/, $ss) {
    for my $w (@w) {
      my $lw = lc($w);
      $s =~ s/$w/$lw/ig;
      $s = reverse $s;
      $s =~ s/$w/$lw/ig;
      $s = reverse $s;
    }
    my $c = ($s =~ y/a-z//);
    $a2 += $c;
  }
}

my $i3 = input(3);
chomp $i3;
my ($w3, $m) = split /\n\n/, $i3;
$w3 =~ s/WORDS://;
my @w3 = split /,/, $w3;
my @l = map {[split //, $_]} split /\n/, $m;
my $h = @l;
my $w = @{$l[0]};

for my $y (0 .. $h - 1) {
  for my $x (0 .. $w - 1) {
  OUTER:
    for my $w3 (@w3) {
      {
        my $m = "";
        for my $i (0 .. (length $w3) - 1) {
          my $nx = ($x + $i) % $w;
          $m .= uc($l[$y]->[$nx]);
        }
        if ($m eq $w3) {
          for my $i (0 .. (length $w3) - 1) {
            my $nx = ($x + $i) % $w;
            $l[$y]->[$nx] = lc($l[$y]->[$nx]);
          }
        }
      }
      {
        my $m = "";
        for my $i (0 .. (length $w3) - 1) {
          my $nx = ($x - $i) % $w;
          $m .= uc($l[$y]->[$nx]);
        }
        if ($m eq $w3) {
          for my $i (0 .. (length $w3) - 1) {
            my $nx = ($x - $i) % $w;
            $l[$y]->[$nx] = lc($l[$y]->[$nx]);
          }
        }
      }
      {
        my $m = "";
        for my $i (0 .. (length $w3) - 1) {
          my $ny = $y + $i;
          last if ($ny >= $h);
          $m .= uc($l[$ny]->[$x]);
        }
        if ($m eq $w3) {
          for my $i (0 .. (length $w3) - 1) {
            my $ny = $y + $i;
            $l[$ny]->[$x] = lc($l[$ny]->[$x]);
          }
        }
      }
      {
        my $m = "";
        for my $i (0 .. (length $w3) - 1) {
          my $ny = $y - $i;
          last if ($ny < 0);
          $m .= uc($l[$ny]->[$x]);
        }
        if ($m eq $w3) {
          for my $i (0 .. (length $w3) - 1) {
            my $ny = $y - $i;
            $l[$ny]->[$x] = lc($l[$ny]->[$x]);
          }
        }
      }
    }
  }
}

my $n = join "", (map {(join "", @$_) . "\n"} @l);
my $a3 = ($n =~ y/a-z//);

print "Part 1: ", $a1, "\n";
print "Part 2: ", $a2, "\n";
print "Part 3: ", $a3, "\n";
