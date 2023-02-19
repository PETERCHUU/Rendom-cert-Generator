Feature: main of UI
  too many of linux to control,  so i want are i site controller

  for security reson, make a login page for after

  Terminal on ui

  so i need are managment software for use.

  want some smooth ui design to easier manage

  scanner for scanning cmd

  add server to server list

  file system for server list

  search engine with file system

  one hit multi Terminal execute design

  preset cmd with multiple Terminal exeute

  store service in speart dictionary

  Scenario: start program (Main)
    Given "env" that have MySQL auth,SSL Generate info
    When startup the execute, first init get the "env" with sql auth
    And Also "env" should have all information that "SSL Generater" need to Generate SSL and TLS
    Then read "env" data to get sql ip,port, AC\PW
    And connect the "sql server"
    But if connect fail
    Then return err and log off
    But if login success
    Then get Web login list from sql to "loginCache"
    And get "TLS Cert" from "SSL Generater"
    And Use "TLS Cert" to Start The HTTPS Api Server




  Scenario: login
    Given json request with login AC\PW
    When server get data from login query
    Then decrpt the data by privite key
    And Read the Post Data by AC and PW
    And Check if AC exist by :"if value, ok := loginCache[{login_Name}];!ok{ //Do somthing}"
    But if login not exist,
    Then respon Login Fail to user
    But if login exist,
    Then encrypt the Pw with md5
    And Compare To Value Of PW In "loginCache"
    And *** if the PW right , Generate the "JWT" from "JWT Generater"


Feature:Security


  Scenario: SSL Generater


  Scenario: JWT Generater


Header with "alg:ES512","typ:JWT" payload with "User_name","Sha256_enerypted_PW","owner_sign"