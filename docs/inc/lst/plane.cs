public class Plane
{
    private IPlaneRepository<Plane> _db;
    [Key]
    public Int64 Id { get; set; }
    public string Manufacturer { get; set; }
    public string Model { get; set; }
    public uint EconomyClassNum{ get; set; }
    public uint BusinessClassNum { get; set; }
    public uint FirstClassNum { get; set; }
}
