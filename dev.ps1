$WORKDIR = Get-Location

switch ($args[0]) {
    "api" {
        Set-Location "$WORKDIR\api"
        go build -o "$WORKDIR\dist\pfila"
        Set-Location "$WORKDIR\runner"
        go build -o "$WORKDIR\dist\pfila_runner.exe"
        Set-Location "$WORKDIR\dist"
        Start-Process -Wait -FilePath "sudo" -ArgumentList "./pfila.exe", "serve"
    }
    "interface" {
        Set-Location "$WORKDIR\interface"
        yarn dev
    }
    "build" {
        Set-Location "$WORKDIR\api"
        go build -o "$WORKDIR\dist\pfila.exe"
        Set-Location "$WORKDIR\runner"
        go build -o "$WORKDIR\dist\pfila_runner.exe"
        Write-Host "Compiling interface..."
        Set-Location "$WORKDIR\interface"
        $env:ENV = "prod"
        npx vite build --emptyOutDir
    }
    default {
        Write-Host "Command unknown"
    }
}
