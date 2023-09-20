CREATE TABLE [TicketsServices] (
	[TicketId] [bigint] CHECK (TicketId > 0),
	[ServiceId] [bigint] CHECK (ServiceId > 0),
	PRIMARY KEY (TicketId, ServiceId),
	CONSTRAINT FK_TICKETS_2_ID
	FOREIGN KEY (TicketId) REFERENCES [Tickets] (Id)
	ON DELETE CASCADE,
	CONSTRAINT FK_SERVICES_ID
	FOREIGN KEY (ServiceId) REFERENCES [Services] (Id)
	ON DELETE CASCADE
);
