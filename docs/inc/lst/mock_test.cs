[Fact]
public void RegisterNewUser()
{
    string email = "testEmail@test.com";
    string password = "testPassword123";
    var mock = new Mock<IUserRepository<User>>();
    mock.Setup(repo => repo.GetUserByEmail(email))
        .Returns(GetNonExistentUserByEmail(email));
    mock.Setup(repo => repo.InsertUser(new User()))
        .Callback(Insert);
    var user = new User(mock.Object);
    user.Register(email, password);
    Assert.Equal(email, user.Email);
    Assert.NotEqual(password, user.Password);
}
