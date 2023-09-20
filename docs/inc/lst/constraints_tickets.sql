CREATE TABLE [Tickets] (
	[Id] [bigint] IDENTITY(1, 1) CHECK (Id > 0),
	[FlightId] [bigint] NOT NULL CHECK (FlightId > 0),
	[OrderId] [bigint],
	[Row] [int] NOT NULL CHECK (Row > 0),
	[Place] [char] NOT NULL,
	[Class] [nvarchar](16) NOT NULL,
	[Refund] [bit] NOT NULL,
	[Price] [money] CHECK (Price > 0),
	PRIMARY KEY (Id),
	CONSTRAINT FK_FLIGHTS_ID
	FOREIGN KEY (FlightId) REFERENCES [Flights] (Id)
	ON DELETE CASCADE,
);
