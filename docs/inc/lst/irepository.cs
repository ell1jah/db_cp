public interface IOrderRepository<T> 
    where T : class
{
    public IEnumerable<T> GetOrderByUserId(Int64 userId);
    public IEnumerable<OrderedTicket> GetTicketsByOrderId(Int64 orderId);
    public T GetActiveOrderByUserId(Int64 userId);
    public void DeleteTicketFromOrder(Int64 ticketId);
    public decimal GetOrderPrice(Int64 orderId);
    public void InsertOrder(Int64 userId);
    public void UpdateOrder(T order);
    public void DeleteOrder(Int64 orderId);
}
