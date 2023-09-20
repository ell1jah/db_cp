CREATE TABLE [Flights] (
	[Id] [bigint] IDENTITY(1, 1) CHECK (Id > 0),
	[PlaneId] [bigint] NOT NULL CHECK (PlaneId > 0),
	[DeparturePoint] [nvarchar](32) NOT NULL,
	[ArrivalPoint] [nvarchar](32) NOT NULL,
	[DepartureDateTime] [datetime] NOT NULL,
	[ArrivalDateTime] [datetime] NOT NULL,
	PRIMARY KEY (Id),
	CHECK (ArrivalDateTime > DepartureDateTime),
	CONSTRAINT FK_PLANES_ID
	FOREIGN KEY (PlaneId) REFERENCES [Planes] (Id)
	ON DELETE CASCADE
);
