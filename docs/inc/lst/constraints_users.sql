CREATE TABLE [Users] (
	[Id] [bigint] IDENTITY(1, 1) CHECK (Id > 0),
	[Role] [varchar](32) NOT NULL,
	[Email] [varchar](32) NOT NULL UNIQUE,
	[Password] [varchar](64) NOT NULL,
	[FirstName] [nvarchar](32),
	[LastName] [nvarchar](32),
	[RegDate] [datetime] NOT NULL,
	PRIMARY KEY (Id),
);
