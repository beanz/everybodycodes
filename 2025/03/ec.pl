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
  return split /,/, $c;
}

my @c = parse(input(1));
@c = uniq sort {$b<=>$a} @c;
print "Part 1: ", (sum @c), "\n";

@c = parse(input(2));
@c = uniq sort {$a<=>$b} @c;
print "Part 2: ", (sum @c[0..19]), "\n";

@c = parse(input(3));
my %m;
for (@c) { $m{$_}++ }
my $c = max values %m;
print "Part 3: ", $c, "\n";

