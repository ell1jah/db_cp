CREATE TABLE [Orders] (
	[Id] [bigint] IDENTITY(1, 1) CHECK (Id > 0),
	[UserId] [bigint] NOT NULL CHECK (UserId > 0),
	[Status] [nvarchar](16) NOT NULL,
	PRIMARY KEY (Id),
	CONSTRAINT FK_USERS_ID
	FOREIGN KEY (UserId) REFERENCES [Users] (Id)
	ON DELETE CASCADE,
);
