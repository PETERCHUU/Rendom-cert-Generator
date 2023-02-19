Get-NetAdapter | Where-Object -Property Status -EQ Up | Select-Object -Property ifIndex
Set-NetIPAddress -InterfaceIndex 4 -IPAddress 10.120.8.168 -AddressFamily IPv4 -PrefixLength 24 -DefaultGateway 10.120.8.1 -
Set-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Services\w32time\Parameters" -Name "NtpServer" -Value time.windows.com, 0x9  
Stop-Service w32time
Start-Service w32time
Restart-Service w32Time
Set-DnsClientServerAddress -InterfaceIndex 4 -ServerAddresses 
Set-NetIPInterface -InterfaceIndex 4