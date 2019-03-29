csv2asv
=======

csv2asv is a processor that converts csv files to use ascii control
separators (Record and Unit). The point of this is to make it easier
to pipe csv files into awk without having to worry about csv quoting
and escaping.

I use the following bash alias to set awk's FS and RS variables appropriately:

    alias awkward='awk -v "FS=\x1F" -v "RS=\x1E"'

Usage example:

    $ echo -e "\"field,with,comma\",\"quote\"\" literal\",val3\\nline2,v2,v3\n"  | ./csv2asv | awkward '{print $2}'
    quote" literal
    v2
