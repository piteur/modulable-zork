#!/usr/bin/env bash
cd bin

cat << EOF
<!DOCTYPE html>
<head>
	<title>Modular Zork - release $TRAVIS_TAG</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
	<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<html lang="en">
	<div class="text-center">
		<h1 class="text-center">Modular Zork</h1>
		<h1><small>Release <strong>#</strong>$TRAVIS_TAG</small></h1>
	</div>

	<div class="row">
		<div class="col-md-offset-1 col-md-10">
			<table class="table table-striped">
				<thead>
					<tr>
						<th>Platform</th>
						<th>386</th>
						<th>amd64</th>
						<th>arm</th>
					</tr>
				</thead>
				<tbody>
EOF

for architecture in */; do
	cd $architecture
	count=0

	echo "<tr><td><strong>${architecture%/}</strong></td>"

	for binary in *; do
		echo "<td><a href="./$architecture$binary">$binary</a></td>"
		let "count++"
	done

	if [ $count -eq 2 ]; then
		echo "<td><i>Not available</i></td>"
	fi

	echo "</tr>"

	cd ..
done

cat << EOF
</tbody>
			</table>
		</div>

		<div class="col-md-offset-3 col-md-6">
			<h3 class="text-center">
				Want to contribute ?
				<small>
					<a href="https://github.com/piteur/modular-zork">
						Come see the code !
					</a>
				</small>
			</h3>

			<pre>$ git clone git@github.com:piteur/modular-zork.git
$ make</pre>

		</div>
	</div>
</html>

EOF
