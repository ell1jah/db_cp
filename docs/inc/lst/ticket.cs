public class Ticket
{
    private ITicketRepository<Ticket> _db;
    [Key]
    public Int64 Id { get; set; }
    public Int64 FlightId { get; set; }
    public Int64 OrderId { get; set; }
    public int Row { get; set; }
    public char Place { get; set; }
    public string Class { get; set; }
    public bool Refund { get; set; }
    public Decimal Price { get; set; }
}
