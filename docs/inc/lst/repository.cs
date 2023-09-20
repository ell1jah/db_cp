public IEnumerable<Order> GetOrderByUserId(Int64 userId)
{
    List<Order> orders = new List<Order>();
    using (SqlConnection connection = new SqlConnection(this._connectionString))
    {
        connection.Open();
        SqlCommand command = new SqlCommand($"SELECT * FROM Orders WHERE UserId = @UserId", connection);
        command.Parameters.AddWithValue("UserId", userId);
        var dataReader = command.ExecuteReader();
        orders = QueryHandler.GetList<Order>(dataReader);
    }
    return orders;
}
