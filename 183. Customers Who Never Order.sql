# Write your MySQL query statement below

SELECT Customers FROM (
    SELECT c.Name AS Customers, o.CustomerId FROM Customers c
        LEFT JOIN Orders o ON
        c.Id = o.CustomerId
    ) orders
    WHERE ISNULL(orders.CustomerId);
