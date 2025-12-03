#include <fcntl.h>
#include <unistd.h>

int	main(int ac, char **av)
{
	if (ac != 2)
		return (write(2, "Program requires one argument\n", 30), 1);
	char input[];
	int	i = 0;
	int	j = 0;
	int	pw = 0;
	int	dial = 50;
	int	fd = open(av[1], O_RDONLY);	
	int	bytes = read(fd, input, sizeof(buff));
	if (bytes < 0)
		return (write(2, "Error reading file\n", 19), 1);
	char	**moves = ft_split(input, '\n');
	if (!moves)
		return (write(2, "Error in split's malloc\n", 24), 1);
	

	while (moves[i])
	{
		int clicks = atoi(&moves[i][1]);
		if (moves[i][0] == 'L')
		{
		
		}
		else if (move[i][0] == 'R')
		{
		
		}
		if (dial == 0)
			pw++;
	}

	return (0)
}
