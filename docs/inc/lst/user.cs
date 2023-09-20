public class User
{
    private IUserRepository<User> _db;
    [Key]
    public Int64 Id { get; set; }
    public string Role { get; set; }
    public string Email { get; set; }
    public string Password { get; set; }
    public string FirstName { get; set; }
    public string LastName { get; set; }
    public DateTime RegDate { get; set; }
}
