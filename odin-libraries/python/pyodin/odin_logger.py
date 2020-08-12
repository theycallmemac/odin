""" Logs stats to Engine """
from requests import post

class OdinLogger: # pylint: disable=too-few-public-methods
    """ OdinLogger is used to access the Engine stats endpoint """

    @classmethod
    def log(cls, stat_type, desc, value, job):
        """ Docstring """
        job_id, job_stamp = job[0], job[1]
        post_data = stat_type +"," + desc + "," + str(value) + "," + job_id + "," + str(job_stamp)
        response = post(url="http://localhost:3939/stats/add", data=post_data)
        return response.status_code
