public class Order
{
    private IOrderRepository<Order> _db;
    [Key]
    public Int64 Id { get; set; }
    public Int64 UserId { get; set; }
    public string Status { get; set; }
}
