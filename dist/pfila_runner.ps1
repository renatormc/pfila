$cmd = "$cmdReplace"
$outputFile = "$outputFileReplace"

$cmd | Out-File -FilePath $outputFile -ErrorAction SilentlyContinue

if ($?) {
    Add-Content -Path $outputFile -Value "PFila: Finish!"
}
