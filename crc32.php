<?php

	$crc = crc32($argv[1]);
	fprintf(STDERR, "CRC of $argv[1] = %08x\n", $crc);
	# crc32("123456789") == d9c60b34
