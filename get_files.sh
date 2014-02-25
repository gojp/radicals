# get radkfile
wget ftp://ftp.monash.edu.au/pub/nihongo/radkfile.gz
gunzip radkfile.gz
iconv -f EUC-JP -t UTF-8 < radkfile > radkfile.utf8

wget ftp://ftp.monash.edu.au/pub/nihongo/kradfile.gz
gunzip kradfile.gz
iconv -f EUC-JP -t UTF-8 < kradfile > kradfile.utf8

rm radkfile kradfile
