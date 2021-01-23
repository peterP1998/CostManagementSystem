# CostManagementSystem
## Overview
Cost Management System is web application that helps you to track your money. You can add incomes and expenses and you can check your balance.
You can help different group targets for example("money for rent" or "money for internet" groups and this groups are crated by your roommate).

<img src="images/expense.PNG" alt="Welcome page">

## Instalation
1. First step is to clone this repo on your pc
```git clone https://github.com/peterP1998/CostManagementSystem.git ```
2. Second step is to install mysql to your pc
3. Then execute db.sql file to create the database
4. To run the application you should type go run main.go. The server will run on port 8090

## Features 
This application have two types of users(admin and regular). Admin user can create groups and users,also can delete users.
Regular user can create expense(for example 10$ for pizza) and create new income(for exmaple 100$ from your parents for your birthday).
Other features of regular user is group donation,you can choose how much you want to donate for group target.Other features is 
expense history and income history. And the last feature is balance,in this page you can check your balance(for example if you spent 10$ for
pizza and you received 100$ from your parents,your balance should be 90$).Also in this page you can view two graphics. Each of these 
graphics are pie charts. For expenses this graphic visualize your expenses by category. For incomes this graphic is similar.

## Final words
This project was created for Golang course in FMI 2021.