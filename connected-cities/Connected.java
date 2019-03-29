import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Map;
import java.util.Queue;
import java.util.Set;

/**
 * This is a simple program to determine if two cities are connected.
 * Connections between cities are listed in a file described below. This file is the first
 * argument. The second and third arguments are the cities which will be checked if there
 * is a path connecting them.
 *
 * Each line of the file indicates that it is possible to travel between the two
 * cities named. (More formally, if we think of the cities as nodes in a graph,
 * each line of the file specifies an edge between two nodes.) In the example
 * above it is possible to travel between Boston and New York, and also between
 * New York and Philadelphia and between Philadelphia and Pittsburgh, so it
 * follows that Boston and Pittsburgh are connected. On the other hand, there is
 * no path between Boston and Tampa, so they are not connected.
 *
 * The file might look something like:
 * Philadelphia, Pittsburgh
 * Boston, New York
 * Philadelphia, New York
 * Los Angeles, San Diego
 * New York, Croton-Harmon
 * St. Petersburg, Tampa
 *
 * <a href="mailto:roland.kluge@gmail.com">Roland Kluge</a>
 */
public class Connected {
	public static final String DELIMITER = ",";

	public static void main(String[] args) {
		if (args == null || args.length != 3) {
			printUsage();
			System.exit(1);
		} else {
			String filename = args[0];
			String city1 = args[1];
			String city2 = args[2];

			try {
				Map<String, Set<String>> cityToNodeMap = parseFileIntoConnections(filename);
				boolean result = isConnected(cityToNodeMap, city1, city2);
				printResult(result);
				System.exit(0);
			} catch (Exception e) {
				String message = e.getMessage();
				System.err.println("Something went wrong. " + message);
				e.printStackTrace();
				System.exit(2);
			}
		}
	}

	private static void printUsage() {
		System.out.println("USAGE: Returns Yes or No, if two cities are connected");
		System.out.println("This program takes 3 arguments");
		System.out.println("\tfilename\ta file containing connected cities, comma separated");
		System.out.println("\tcityname1\tthe first city");
		System.out.println("\tcityname2\tthe second city");
	}

	private static void printResult(boolean result) {
		if (result) {
			System.out.println("yes");
		} else {
			System.out.println("no");
		}
	}

	// The key in the map are the cities that we have available.
	// The value is a set of string, a set of cities that this city is connected to
	// So the strings are the nodes, the name of the cities.
	// If there is a key-value pair (or key to set pair) then those cities are connected
	//
	// The map approach is sparse, so it will reduce the amount of memory used
	// however, for each line, we add two key-value pairs for both directions of possible traversal
	private static Map<String, Set<String>> parseFileIntoConnections(String filename) throws IOException {
		Map<String, Set<String>> cityToNodeMap = new HashMap<String, Set<String>>();

		BufferedReader bufferedReader = null;
		try {
			Reader fileReader = new FileReader(filename);
			bufferedReader = new BufferedReader(fileReader);
			String line = bufferedReader.readLine();
			while (line != null && !line.isEmpty()) {
				String[] cities = line.split(DELIMITER);
				String firstCity = cities[0].trim();
				String secondCity = cities[1].trim();

				Set<String> firstCityConnections = getCityConnections(cityToNodeMap, firstCity);
				Set<String> secondCityConnections = getCityConnections(cityToNodeMap, secondCity);
				firstCityConnections.add(secondCity);
				secondCityConnections.add(firstCity);

				line = bufferedReader.readLine();
			}
		} finally {
			if (bufferedReader != null) {
				bufferedReader.close();
			}
		}

		return cityToNodeMap;
	}

	// helper function to default the connected cities in the map to the empty set
	private static Set<String> getCityConnections(Map<String, Set<String>> map, String city) {
		if (!map.containsKey(city)) {
			map.put(city, new HashSet<String>());
		}
		return map.get(city);
	}


	private static boolean isConnected(Map<String, Set<String>> cityToNodeMap, String city1, String city2) {
		boolean isFound = city1.equals(city2);
		if (cityToNodeMap.containsKey(city1) && cityToNodeMap.containsKey(city2)) {
			// By using a Queue, we are implementing Breadth First Search
			// This will find the shortest path between two cities
			Queue<String> citiesToVisit = new LinkedList<String>();

			// We keep a set of the cities we have already visited. This prevents BFS from looping in
			// cycles and allows the BFS to terminate if no path can be found after exploring all reachable nodes
			Set<String> citiesAlreadyVisited = new HashSet<String>();

			citiesToVisit.add(city1);

			while (!citiesToVisit.isEmpty() && !isFound) {
				String city = citiesToVisit.poll();
				isFound = city.equals(city2);

				Set<String> possibleConnections = cityToNodeMap.get(city);
				for (String possibleCity : possibleConnections) {
					if (!citiesAlreadyVisited.contains(possibleCity)) {
						citiesToVisit.add(possibleCity);
						citiesAlreadyVisited.add(possibleCity);
					}
				}
			}
		}

		return isFound;
	}
}
