import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

public class DataSetManager {
	private Map<Integer, Integer> dataSetToCenterMap = new HashMap<Integer, Integer>();
	private List<DataCenter> dataCenterList;

	public DataSetManager(int total) {
		dataCenterList = new ArrayList<DataCenter>(total);
	}

	public void addDataCenter(int id, String line) {
		String[] tokens = line.split("\\s");
		DataCenter dataCenter = new DataCenter(id);

		for(String s: tokens) {
			int dataSetId = Integer.valueOf(s);
			dataCenter.addDataSet(dataSetId);
			dataCenterList.add(dataCenter);
			dataSetToCenterMap.put(dataSetId, id);
		}
	}

	public Set<Copy> getCopies() {
		Set<Copy> results = new HashSet<Copy>();

		for(DataCenter dataCenter: dataCenterList) {
			int dataCenterDest = dataCenter.getId();
			Set<Integer> missingDataSets = dataCenter.getMissingDataSets();
			for(Integer data: missingDataSets) {
				int dataCenterSource = dataSetToCenterMap.get(data); // copy it from dataCenterSource
				Copy copy = new Copy(dataCenterSource, dataCenterDest, data);
				results.add(copy);
			}
		}

		return results;
	}
}
