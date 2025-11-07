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
  return split /\n/, $c;
}

my @c = parse(input(1));
my $a1 = floor(2025*$c[0]/$c[$#c]);
print "Part 1: ", $a1, "\n";

@c = parse(input(2));
my $a2 = ceil(10000000000000*$c[$#c]/$c[0]);
print "Part 2: ", $a2, "\n";

@c = parse(input(3));
my $s = shift @c;
my $l = pop @c;
for (@c) {
  my @n = split /\|/;
  $s /= $n[0];
  $s *= $n[1];
}
my $a3 = ceil(100*$s/$l);
print "Part 3: ", $a3, "\n";
