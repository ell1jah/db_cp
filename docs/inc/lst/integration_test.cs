[Fact]
public void RegisterNewUser()
{
    string email = "testEmail@test.com";
    string password = "testPassword123";
    var user = new User();
    user.Register(email, password);
    Assert.Equal(email, user.Email);
    Assert.NotEqual(password, user.Password);
}
