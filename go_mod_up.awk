/^replace \(/, /^\)/{
	if($0~/^\s+/)
		system("go get -u "$3)
}
/^require \(/, /^\)/{
	if($0~/^\s+/){
	    if($0~/caddy/ || $0~/fiber/ || $0~/rpcx/ || $0~/sno/)
		    system("go get -u "$1"@master")
		else if($0~/grpc/)
            system("go get -u "$1"@v1.26.0")
        else
		    system("go get -u "$1)
	}
}
