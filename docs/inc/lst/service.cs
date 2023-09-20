public class Service
{
    private IServiceRepository<Service> _db;
    [Key]
    public Int64 Id { get; set; }
    public string Name { get; set; }
    public Decimal Price { get; set; }
    public bool EconomyClass { get; set; }
    public bool BusinessClass { get; set; }
    public bool FirstClass { get; set; }
}
