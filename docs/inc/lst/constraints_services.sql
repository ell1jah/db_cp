CREATE TABLE [Services] (
	[Id] [bigint] IDENTITY(1, 1) CHECK (Id > 0),
	[Name] [nvarchar](64) NOT NULL,
	[Price] [money] NOT NULL CHECK (Price > 0),
	[EconomyClass] [bit] NOT NULL,
	[BusinessClass] [bit] NOT NULL,
	[FirstClass] [bit] NOT NULL,
	PRIMARY KEY (Id),
);
