# Vending Machine Design Problem

## Introduction
As part of this interview, I would like you to design a vending machine system. This will help us assess your problem-solving skills, understanding of object-oriented design principles, and ability to think through system requirements.

## Problem Statement
Imagine we want to create a vending machine that dispenses snacks and beverages. The vending machine should be able to handle various operations, such as:

1. **Displaying Items**: Show available items with their prices.
2. **Accepting Payments**: Accept coins and bills, and provide change.
3. **Dispensing Items**: Dispense the selected item after payment is confirmed.
4. **Restocking**: Allow for restocking of items.
5. **Inventory Management**: Keep track of the quantity of each item and notify when items are low in stock.
6. **User Interface**: Provide a simple user interface for interaction.

## Requirements

### Functional Requirements
- The machine should allow users to select an item by entering its code.
- The machine should accept multiple forms of payment (coins, bills).
- It should return change if the payment exceeds the item price.
- The system should handle inventory checks before dispensing an item.
- An admin interface should be available for restocking items.

### Non-Functional Requirements
- The system should be reliable and handle concurrent transactions.
- It should be easy to maintain and extend for future enhancements.

## Questions for You
1. **Class Design**: Can you outline the main classes you would create for this vending machine? What attributes and methods would they have?

2. **State Management**: How would you manage the state of the vending machine (e.g., available items, current balance, etc.)?

3. **Payment Handling**: What approach would you take to handle payments? How would you ensure that the correct change is provided?

4. **Error Handling**: How would you handle errors, such as insufficient funds or out-of-stock items?

5. **Extensibility**: If we wanted to add new features in the future (like loyalty points or a mobile app interface), how would you design the system to accommodate these changes?

## Conclusion
Please take your time to think through these aspects and share your design approach with us. Feel free to use diagrams or pseudocode if that helps illustrate your thought process!