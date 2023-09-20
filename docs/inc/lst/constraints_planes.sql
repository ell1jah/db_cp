CREATE TABLE [Planes] (
	[Id] [bigint] IDENTITY(1, 1) CHECK (Id > 0),
	[Manufacturer] [nvarchar](32) NOT NULL,
	[Model] [nvarchar](16) NOT NULL,
	[EconomyClassNum] [integer] NOT NULL CHECK (EconomyClassNum >= 0),
	[BusinessClassNum] [integer] NOT NULL CHECK (BusinessClassNum >= 0),
	[FirstClassNum] [integer] NOT NULL CHECK (FirstClassNum >= 0),
	PRIMARY KEY (Id),
);
