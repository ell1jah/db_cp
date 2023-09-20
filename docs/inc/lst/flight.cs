public class Flight
{
    private IFlightRepository<Flight> _db;
    [Key]
    public Int64 Id { get; set; }
    public Int64 PlaneId { get; set; }
    public string DeparturePoint { get; set; }
    public string ArrivalPoint { get; set; }
    public DateTime DepartureDateTime { get; set; }
    public DateTime ArrivalDateTime { get; set; }
}
