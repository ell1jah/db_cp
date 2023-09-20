CREATE FUNCTION GetOrderPrice (@OrderId bigint)
RETURNS money
BEGIN
    DECLARE @sumTickets money, @sumServices money, @result money
    SET @sumTickets = 0;
    SET @sumServices = 0;
    SET @result = 0;
    SELECT @sumTickets = COALESCE(SUM(t.Price), 0) FROM Orders o 
    JOIN Tickets t ON o.Id = t.OrderId WHERE o.Id = @OrderId;
    SELECT @sumServices = COALESCE(SUM(s.Price), 0) FROM Orders o 
    JOIN Tickets t ON t.OrderId = @OrderId JOIN
    TicketsServices ts ON t.Id = ts.TicketId 
    JOIN Services s ON s.Id = ts.ServiceId WHERE o.Id = @OrderId;
    SELECT @result = @sumTickets + @sumServices
RETURN @result
END
